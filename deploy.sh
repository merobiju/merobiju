#!/bin/bash
set -e

echo "🚀 Generating README via Go..."
go run generate_readme.go

echo "📦 Staging changes..."
git add .

echo "💾 Committing..."
git commit -m "Auto-update: Generated README with Go script" || echo "ℹ️ No changes to commit."

echo "☁️ Pushing to GitHub..."
git push origin main

echo "✅ All done! Your profile is updated."
