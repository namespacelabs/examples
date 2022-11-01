// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

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
