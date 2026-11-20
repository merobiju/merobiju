package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cacheBuster := time.Now().Unix()

	readmeContent := fmt.Sprintf(`<div align="center">
  <img src="https://raw.githubusercontent.com/merobiju/merobiju/main/assets/banner.png" alt="Digital Sadhak - Bijaya Pokhrel" width="100%%" />
  
  # Hi, I'm Bijaya Pokhrel (Digital Sadhak) <img src="https://upload.wikimedia.org/wikipedia/commons/e/e7/Animated-Flag-Nepal.gif" width="45" height="30" />
  
  <a href="https://git.io/typing-svg"><img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=700&size=28&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=Self-Hosted+Cloud+Architect;Linux+Tinkerer;Go+Developer;Open-Source+Advocate" alt="Typing SVG" /></a>

  *"मानव इतिहास र अनन्त भविष्यको सेतु बन्दै, मेसिनको युगमा मानवता र आत्मीयता जोगाउने एक डिजिटल साधक ✨ღ♾️ 🇳🇵 | Self-Hosted & Go Developer"*
  
  <a href="https://www.google.com/maps/place/Kathmandu,+Nepal" target="_blank"><img src="https://img.shields.io/badge/Kathmandu%%2C_Nepal-0052CC?style=for-the-badge&logo=google-maps&logoColor=white" alt="Location" /></a>
  <a href="https://www.cloudflare.com/zero-trust/" target="_blank"><img src="https://img.shields.io/badge/Security-Zero--Trust-CC0000?style=for-the-badge&logo=cloudflare&logoColor=white" alt="Zero Trust" /></a>
  <a href="https://github.com/merobiju" target="_blank"><img src="https://komarev.com/ghpvc/?username=merobiju&style=for-the-badge&color=brightgreen" alt="Profile Views" /></a>
</div>

---

### 🛠️ The Digital Fortress: Ultimate Tech Stack
I engineer and maintain a sovereign, private cloud ecosystem from scratch, ensuring absolute data privacy and digital independence.

<div align="center">

  <!-- 1. Hardware, OS & Compute -->
  <h4>🖥️ Hardware, OS & Compute</h4>
  <a href="https://www.amd.com/" target="_blank"><img src="https://cdn.jsdelivr.net/npm/@lobehub/icons-static-webp@latest/light/amdradeoncloud-color.webp" alt="AMD" height="45" title="AMD Ryzen" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.nvidia.com/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/nvidia.webp" alt="NVIDIA" height="45" title="NVIDIA" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://ubuntu.com/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/ubuntu.webp" alt="Ubuntu" height="45" title="Ubuntu" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.kernel.org/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/linux.webp" alt="Linux" height="45" title="Linux" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.debian.org/" target="_blank"><img src="https://cdn.simpleicons.org/debian/A81D33" alt="Debian" height="45" title="Debian" /></a>
  <br/><br/>

  <!-- 2. Infrastructure, Networking & Security -->
  <h4>🌐 Infrastructure, Networking & Security</h4>
  <a href="https://www.microsoft.com/windows" target="_blank"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/windows8/windows8-original.svg" alt="Windows" height="45" title="Windows" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.docker.com/" target="_blank"><img src="https://cdn.simpleicons.org/docker/2496ED" alt="Docker" height="45" title="Docker" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://kubernetes.io/" target="_blank"><img src="https://cdn.simpleicons.org/kubernetes/326CE5" alt="Kubernetes" height="45" title="Kubernetes" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.portainer.io/" target="_blank"><img src="https://cdn.simpleicons.org/portainer/13BEF9" alt="Portainer" height="45" title="Portainer" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://tailscale.com/" target="_blank"><img src="https://cdn.simpleicons.org/tailscale/242424" alt="Tailscale" height="45" title="Tailscale" /></a>
  <br/><br/>
  
  <a href="https://nextdns.io/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/nextdns.webp" alt="NextDNS" height="45" title="NextDNS" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.cloudflare.com/zero-trust/" target="_blank"><img src="https://cdn.simpleicons.org/cloudflare/F38020" alt="Cloudflare" height="45" title="Cloudflare" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://brave.com/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/brave.webp" alt="Brave" height="45" title="Brave" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://github.com/degoog-org/degoog" target="_blank"><img src="https://raw.githubusercontent.com/homarr-labs/dashboard-icons/refs/heads/main/webp/degoog.webp" alt="DeGoog" height="45" title="DeGoog" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://github.com/huashengdun/webssh" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/webssh.webp" alt="WebSSH" height="45" title="WebSSH" /></a>
  <br/><br/>

  <a href="https://microsoft.com/powershell" target="_blank"><img src="https://raw.githubusercontent.com/homarr-labs/dashboard-icons/main/webp/powershell.webp" alt="PowerShell" height="45" title="PowerShell" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://nginx.org/" target="_blank"><img src="https://cdn.simpleicons.org/nginx/009639" alt="Nginx" height="45" title="Nginx" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.wireguard.com/" target="_blank"><img src="https://cdn.simpleicons.org/wireguard/88171A" alt="WireGuard" height="45" title="WireGuard" /></a>
  <br/><br/>

  <!-- 3. Development, Languages & Automation -->
  <h4>💻 Development, Languages & Automation</h4>
  <a href="https://go.dev/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/webp/go.webp" alt="Go" height="35" title="Go" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.python.org/" target="_blank"><img src="https://cdn.simpleicons.org/python/3776AB" alt="Python" height="45" title="Python" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.gnu.org/software/bash/" target="_blank"><img src="https://cdn.simpleicons.org/gnubash/4EAA25" alt="Bash" height="45" title="Bash" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://git-scm.com/" target="_blank"><img src="https://cdn.simpleicons.org/git/F05032" alt="Git" height="45" title="Git" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://github.com/" target="_blank"><img src="https://cdn.simpleicons.org/github/181717" alt="GitHub" height="45" title="GitHub" /></a>
  <br/><br/>
  
  <a href="https://github.com/features/actions" target="_blank"><img src="https://cdn.simpleicons.org/githubactions/2088FF" alt="GitHub Actions" height="45" title="CI/CD" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://code.visualstudio.com/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/vscode/vscode-original.svg" alt="VS Code" height="45" title="VS Code" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://neovim.io/" target="_blank"><img src="https://cdn.simpleicons.org/neovim/57A143" alt="Neovim" height="45" title="Neovim" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.json.org/" target="_blank"><img src="https://cdn.simpleicons.org/json/000000" alt="JSON" height="45" title="JSON" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://daringfireball.net/projects/markdown/" target="_blank"><img src="https://cdn.simpleicons.org/markdown/000000" alt="Markdown" height="45" title="Markdown" /></a>
  <br/><br/>

  <!-- 4. Data, Databases & Monitoring -->
  <h4>📊 Data, Databases & Monitoring</h4>
  <a href="https://www.postgresql.org/" target="_blank"><img src="https://cdn.simpleicons.org/postgresql/4169E1" alt="PostgreSQL" height="45" title="PostgreSQL" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://redis.io/" target="_blank"><img src="https://cdn.simpleicons.org/redis/DC382D" alt="Redis" height="45" title="Redis" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.sqlite.org/" target="_blank"><img src="https://cdn.simpleicons.org/sqlite/003B57" alt="SQLite" height="45" title="SQLite" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://min.io/" target="_blank"><img src="https://cdn.simpleicons.org/minio/C72E49" alt="MinIO" height="45" title="MinIO" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://prometheus.io/" target="_blank"><img src="https://cdn.simpleicons.org/prometheus/E6522C" alt="Prometheus" height="45" title="Prometheus" /></a>
  <br/><br/>
  
  <a href="https://grafana.com/" target="_blank"><img src="https://cdn.simpleicons.org/grafana/F46800" alt="Grafana" height="45" title="Grafana" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://uptime.kuma.pet/" target="_blank"><img src="https://raw.githubusercontent.com/louislam/uptime-kuma/master/public/icon.svg" alt="Uptime Kuma" height="45" title="Uptime Kuma" /></a>
  <br/><br/>

  <!-- 5. Self-Hosted, Media & Apps -->
  <h4>📸 Self-Hosted, Media & Apps</h4>
  <a href="https://immich.app/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/immich.webp" alt="Immich" height="45" title="Immich" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://proton.me/mail" target="_blank"><img src="https://cdn.jsdelivr.net/gh/selfhst/icons/webp/proton.webp" alt="ProtonMail" height="45" title="ProtonMail" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.videolan.org/vlc/" target="_blank"><img src="https://cdn.simpleicons.org/vlcmediaplayer/FF8800" alt="VLC" height="45" title="VLC" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.craft.do/" target="_blank"><img src="https://cdn.jsdelivr.net/gh/homarr-labs/dashboard-icons/png/craft-app.png" alt="Craft" height="45" title="Craft" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://jellyfin.org/" target="_blank"><img src="https://cdn.simpleicons.org/jellyfin/00A4DC" alt="Jellyfin" height="45" title="Jellyfin" /></a>
  <br/><br/>
  
  <a href="https://www.plex.tv/" target="_blank"><img src="https://cdn.simpleicons.org/plex/E5A00D" alt="Plex" height="45" title="Plex" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.home-assistant.io/" target="_blank"><img src="https://cdn.simpleicons.org/homeassistant/18BCF2" alt="Home Assistant" height="45" title="Home Assistant" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://ffmpeg.org/" target="_blank"><img src="https://upload.wikimedia.org/wikipedia/commons/5/5f/FFmpeg_Logo_new.svg" alt="FFmpeg" height="35" title="FFmpeg" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://www.wikipedia.org/" target="_blank"><img src="https://cdn.simpleicons.org/wikipedia/000000" alt="Wikipedia" height="45" title="Wikipedia" /></a>
  
</div>

<br />

- **Infrastructure:** Ubuntu 26.04 LTS, Docker & Compose, Custom Bare-Metal Home Server (AMD Ryzen 3).
- **Networking & Security:** Tailscale Mesh VPN, NextDNS, Cloudflare Zero-Trust, WebSSH, PowerShell.
- **Media & Data:** FFmpeg with NVIDIA NVENC, Immich, ProtonMail, PostgreSQL, Redis.

---

### 💡 Engineering Philosophy & Pillars
- 🔒 **Digital Sovereignty:** Owning my data and infrastructure, not renting it. 
- 🌐 **Continuous Learning:** Exploring new skills, solving real-world problems, and sharing knowledge.
- 📖 **Nepali Wikipedia Contributor:** Bridging the knowledge gap in the open-source ecosystem.
- 🌿 **Nature & Balance:** Trekking, literature, and music fuel the "Digital Sadhak" journey.

---

### 🚧 Currently Working On
- 🛠️ Architecting **immich-pulse v1.1.0** with advanced telemetry and metrics.
- 🌐 Expanding my self-hosted Zero-Trust network (Tailscale + Docker).
- 📖 Contributing to the Nepali Wikipedia open knowledge ecosystem.

---

### 🚀 Flagship Project: Go on the Edge
- **[immich-pulse](https://github.com/merobiju/immich-pulse)** 
  *An ultra-lightweight (~3.8MB) static Go TUI dashboard built with Charmbracelet's Bubble Tea.*
  *   🖥️ **Running Live** on my Ubuntu Home Server (` + "`merobiju`" + `).
  *   📊 Monitors Immich Postgres stats, Docker container health, CPU/Temp, and Network I/O in real-time.
  *   🔒 Includes a 1-Key Privacy Mode (` + "`p`" + `) to mask IPs instantly.
  *   🌍 Cross-compiled for 14 OS/Architectures from a single ` + "`CGO_ENABLED=0`" + ` binary.
- **[immich-deb-installer](https://github.com/merobiju/immich-deb-installer)** - Automated Debian/Ubuntu deployment script for Immich.

### 🖥️ Home Server Specs (The Digital Fortress)
- **Compute:** AMD Ryzen 3 3250U (4 Threads) | 9.2GB RAM | 467GB NVMe SSD
- **OS:** Ubuntu 26.04.1 LTS (Kernel 7.0.0-44)
- **Stack:** Docker & Compose, Tailscale Mesh VPN, FFmpeg (NVENC), Immich
- **Hardware Guard:** Battery charge threshold locked at 80%% for 24/7 longevity.

---

### 📊 GitHub Telemetry
<div align="center">
  <img src="https://github-readme-stats.vercel.app/api?username=merobiju&show_icons=true&theme=radical&hide_border=true&count_private=true&hide=stars,issues&v=%d" alt="Bijaya's GitHub Stats" height="180" />
  <img src="https://github-readme-stats.vercel.app/api/top-langs/?username=merobiju&layout=compact&theme=radical&hide_border=true&hide=shell,html,css&v=%d" alt="Top Languages" height="180" />
</div>

<div align="center">
  <img src="https://github-readme-activity-graph.vercel.app/graph?username=merobiju&theme=react-dark&hide_border=true&area=true&v=%d" width="100%%" />
</div>

---

### 🌐 Connect with Me
<div align="center">
  <a href="https://ne.wikipedia.org/s/18kv" target="_blank"><img src="https://img.shields.io/badge/Wikipedia-000000?style=for-the-badge&logo=wikipedia&logoColor=white" /></a>
  <a href="https://www.facebook.com/ItsBijaya/" target="_blank"><img src="https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white" /></a>
  <a href="https://x.com/merobiju" target="_blank"><img src="https://img.shields.io/badge/X-000000?style=for-the-badge&logo=x&logoColor=white" /></a>
  <a href="mailto:merobiju@proton.me" target="_blank"><img src="https://img.shields.io/badge/ProtonMail-8B89CC?style=for-the-badge&logo=protonmail&logoColor=white" /></a>
</div>

<br />

<div align="center">
  <i>"A continuous journey of nature, knowledge, technology, and open source."</i><br>
  <b>Let's connect and build a sovereign, open, and digital future!</b>
</div>

<img src="https://capsule-render.vercel.app/api?type=waving&color=00ADD8&height=100&section=footer" width="100%%" />
`, cacheBuster, cacheBuster, cacheBuster)

	err := os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write README.md: %v", err)
	}

	fmt.Printf("✅ README.md generated with Nepali quote and Animated Flag! Cache: %d\n", cacheBuster)
}
