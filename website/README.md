# TermNote Website

Static website showcasing the TermNote terminal note-taking application.

## Tech Stack

- Astro 4.0
- Tailwind CSS
- TypeScript

## Development

### Prerequisites

- Node.js 18.20.8 or higher
- npm or yarn

### Setup

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## Project Structure

```
website/
├── src/
│   ├── components/      # Reusable components
│   │   ├── Hero.astro
│   │   ├── Features.astro
│   │   ├── Installation.astro
│   │   ├── Shortcuts.astro
│   │   └── Footer.astro
│   ├── layouts/         # Page layouts
│   │   └── BaseLayout.astro
│   └── pages/           # Routes
│       └── index.astro
├── public/              # Static assets
├── astro.config.mjs     # Astro configuration
├── tailwind.config.mjs  # Tailwind configuration
└── package.json
```

## Deployment

The website is a static site and can be deployed to:

- Vercel
- Netlify
- GitHub Pages
- Cloudflare Pages
- Any static hosting service

### Build Command

```bash
npm run build
```

### Output Directory

```
dist/
```
