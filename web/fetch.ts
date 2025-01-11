export async function mahlzeitFetch<T>(
  path: string,
  body?: BodyInit
): Promise<T> {
  return new Promise(async (res, rej) => {
    const config = useRuntimeConfig();
    let response = await fetch(config.public.server + path, {
      body: body,
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (response.status > 299) {
      rej(`got non 200 code: ${response.status}`);
    }
    res((await response.json())["data"] as T);
  });
}
