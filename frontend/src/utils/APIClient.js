
import {
    LOCAL_API_URL,
    SPONSORS_URL
  } from './Constants'

const APICall = (url, headers, convertToJson = true) => {
    return fetch(url, headers)
            .then(res => convertToJson ? res.json() : res)
            .then((res) => {
            return res
        })
}

const getHeaders = (token) => {
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    }
    if (token) headers.Authorization = 'Bearer ' + token
    return headers
}

const getClient = (method, body) => {
    const client = {
        method: method,
        headers: getHeaders(),
    }
    if (body) client.body = body
    return client;
}

const sponsorsAPI = () => {
    const url = LOCAL_API_URL + SPONSORS_URL
    const client = getClient('GET')
    return APICall(url, client)
}

const APIClient = {
    sponsorsAPI
}
export default APIClient