import qs from "qs"
import * as fs from "fs"
import { Backends } from "../src/config/backends.ns"

/**
 * Get full Strapi URL from path
 * @param {string} path Path of the URL
 * @returns {string} Full Strapi URL
 */
export function getStrapiURL(path = "") {
  if (!backendUrl) {
    backendUrl = backendUrlFromNsConfig()
  }

  return `${backendUrl}${path}`
}

/**
 * Helper to make GET requests to Strapi API endpoints
 * @param {string} path Path of the API route
 * @param {Object} urlParamsObject URL params object, will be stringified
 * @param {Object} options Options passed to fetch
 * @returns Parsed API call response
 */
export async function fetchAPI(path, urlParamsObject = {}, options = {}) {
  // Merge default and user options
  const mergedOptions = {
    headers: {
      "Content-Type": "application/json",
    },
    ...options,
  }

  // Build request URL
  const queryString = qs.stringify(urlParamsObject)
  const requestUrl = `${getStrapiURL(
    `/api${path}${queryString ? `?${queryString}` : ""}`
  )}`

  // Trigger API call
  const response = await fetch(requestUrl, mergedOptions)

  // Handle response
  if (!response.ok) {
    console.error(response.statusText)
    throw new Error(`An error occured please try again`)
  }
  const data = await response.json()
  return data
}

let backendUrl

function backendUrlFromNsConfig() {
  if (typeof window !== "undefined") {
    // Browser: using a public/localhost address
    return Backends.strapibackend.managed
  } else {
    // Server: using an in-cluster address
    const nsConfigRaw = fs.readFileSync("/namespace/config/runtime.json")
    const nsConfig = JSON.parse(nsConfigRaw.toString())

    const backendService = nsConfig.stack_entry
      .map((e) => e.service)
      .flat()
      .find((s) => s.name === "backendapi")

    return `http://${backendService.endpoint}`
  }
}
