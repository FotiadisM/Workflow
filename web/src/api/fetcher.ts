import { serverURI } from "./url";

export async function fetcher<T>(
  url: string,
  access_token?: string
): Promise<T> {
  const res = await fetch(serverURI + url, {
    headers: {
      Authorization: "Bearer " + access_token,
    },
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(res.statusText + " " + text);
  }

  return res.json() as Promise<T>;
}
