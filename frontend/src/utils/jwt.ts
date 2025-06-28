export function parseJwt(token: string): any | null {
  try {
    const base64 = token.split('.')[1];
    const json = atob(base64);
    return JSON.parse(json);
  } catch {
    return null;
  }
}
