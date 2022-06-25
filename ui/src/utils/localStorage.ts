const accessTokenKey = 'grAccessToken'

export const setAccessToken = (accessToken: string) => {
  localStorage.setItem(accessTokenKey, accessToken)
}

export const getAccessToken = () : string | null => {
  return localStorage.getItem(accessTokenKey)
}
