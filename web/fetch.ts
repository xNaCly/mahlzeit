const SERVER = "http://localhost:8080/api/";

export async function mahlzeitFetch<T>(
  path: string,
  body?: BodyInit
): Promise<T> {
  return new Promise(async (res, rej) => {
    let response = await fetch(SERVER + path, {
      body: body,
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (response.status > 299) {
      rej(`got non 200 code: ${response.status}`);
    }
    res(await response.json());
  });
}
