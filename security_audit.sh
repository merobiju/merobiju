#!/bin/bash
# ==========================================
# Ubuntu Server Security Audit Script
# For: merobiju (Digital Sadhak)
# ==========================================
set -e

echo "🔍 Starting Security Audit on $(hostname) at $(date)"
echo "=================================================="

# 1. System Information
echo -e "\n[1] System Information:"
uname -a
lsb_release -a 2>/dev/null || cat /etc/os-release

# 2. Open Ports (Listening Services)
echo -e "\n[2] Open Ports (Listening Services):"
sudo ss -tulpn | grep -E 'LISTEN|UNCONN' || echo "No listening ports found."

# 3. Firewall Status (UFW & iptables)
echo -e "\n[3] Firewall Status:"
if command -v ufw &> /dev/null; then
    sudo ufw status verbose
else
    echo "UFW not installed."
fi
echo "--- iptables rules ---"
sudo iptables -L -n -v 2>/dev/null | head -20 || echo "iptables not accessible."

# 4. SSH Configuration Audit
echo -e "\n[4] SSH Configuration Audit:"
if [ -f /etc/ssh/sshd_config ]; then
    echo "--- Key SSH Settings ---"
    sudo sshd -T 2>/dev/null | grep -iE 'permitrootlogin|passwordauthentication|x11forwarding|maxauthtries|logingracetime|allowusers|port' || echo "Could not read effective SSH config."
else
    echo "SSH server not installed."
fi

# 5. Failed Login Attempts
echo -e "\n[5] Recent Failed Login Attempts:"
sudo lastb -n 5 2>/dev/null || echo "No failed login records (or lastb not available)."

# 6. Docker Security Check
echo -e "\n[6] Docker Security Check:"
if command -v docker &> /dev/null; then
    echo "--- Running Containers ---"
    docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}"
    echo "--- Docker Socket Permissions ---"
    ls -l /var/run/docker.sock 2>/dev/null || echo "Docker socket not found."
    echo "--- Containers Running as Root ---"
    docker ps -q | xargs -r docker inspect --format='{{.Name}} -> User: {{.Config.User}}' | grep -E 'User: root|User: $' || echo "All containers have non-root user specified (or user is empty)."
else
    echo "Docker not installed."
fi

# 7. Tailscale Status & Peers
echo -e "\n[7] Tailscale Status:"
if command -v tailscale &> /dev/null; then
    tailscale status
    echo "--- Tailscale Peers (JSON summary) ---"
    tailscale status --json | jq -r '.Peer[] | "\(.HostName) (\(.OS)) - Online: \(.Online)"' 2>/dev/null || tailscale status --json | grep -E 'HostName|Online' || echo "Could not parse Tailscale JSON."
else
    echo "Tailscale not installed."
fi

# 8. Sudoers & User Privileges
echo -e "\n[8] Sudoers & User Privileges:"
sudo grep -E '^[^#].*ALL=(ALL)' /etc/sudoers /etc/sudoers.d/* 2>/dev/null || echo "No sudoers entries found (or no permission)."
echo "--- Users with UID 0 (root equivalent) ---"
awk -F: '$3 == 0 {print $1}' /etc/passwd

# 9. File Permissions (World-Writable)
echo -e "\n[9] World-Writable Files (excluding /proc, /sys, /dev, /tmp):"
sudo find / -xdev -type f -perm -0002 ! -path "/proc/*" ! -path "/sys/*" ! -path "/dev/*" ! -path "/tmp/*" 2>/dev/null | head -10 || echo "No world-writable files found (or find timed out)."

# 10. SUID/SGID Binaries
echo -e "\n[10] SUID/SGID Binaries (recent or unexpected):"
sudo find / -xdev \( -perm -4000 -o -perm -2000 \) -type f 2>/dev/null | head -15 || echo "Could not list SUID/SGID binaries."

# 11. Kernel & Package Updates
echo -e "\n[11] Pending Security Updates:"
sudo apt update -qq 2>/dev/null && sudo apt list --upgradable 2>/dev/null | grep -i security || echo "No security updates pending (or apt not configured)."

# 12. Lynis Installation Suggestion
echo -e "\n[12] Deep Security Audit (Lynis):"
if command -v lynis &> /dev/null; then
    echo "Lynis is installed. Run: sudo lynis audit system"
else
    echo "Lynis not installed. Install it for a deeper audit:"
    echo "  sudo apt install lynis"
    echo "  sudo lynis audit system"
fi

echo -e "\n=================================================="
echo "✅ Security Audit Complete."
