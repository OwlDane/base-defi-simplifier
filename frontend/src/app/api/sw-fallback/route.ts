import { NextResponse } from 'next/server'

export async function GET() {
  const swContent = `
    // Development fallback service worker
    self.addEventListener('install', () => {
    console.log('Service worker installed (development fallback)');
    });

    self.addEventListener('activate', () => {
    console.log('Service worker activated (development fallback)');
    });
  `.trim()

  return new NextResponse(swContent, {
    headers: {
      'Content-Type': 'application/javascript',
      'Cache-Control': 'no-cache, no-store, must-revalidate',
    },
  })
}