import { Backends } from "../src/config/backends.ns"

export function getStrapiMedia(media) {
  const { url } = media.data.attributes
  const imageUrl = url.startsWith("/") ? getStrapiMediaURL(url) : url
  return imageUrl
}

// The media is always fetched from the browser, so need to use the browser Strapi URL for that.
function getStrapiMediaURL(path = "") {
  return `${Backends.strapibackend.managed}${path}`
}
