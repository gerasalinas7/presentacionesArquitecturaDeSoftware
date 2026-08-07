# Presentaciones de Arquitectura de Software

Sitio estático para publicar clases online con GitHub Pages.

## Estructura

- `site/`: sitio publicado
- `site/clases/`: una carpeta por clase
- `.github/workflows/deploy.yml`: deploy automático a GitHub Pages

## Clase 1

La primera clase quedó migrada a una presentación local en Reveal.js usando el
contenido extraído del PDF.

## Cómo seguir

- duplicar `site/clases/clase-1/index.html` como base para nuevas clases
- mantener recursos visuales compartidos en `site/assets/`
- el deploy publica `site/` directo en GitHub Pages, sin build step
