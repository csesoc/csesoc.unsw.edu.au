
import {
    BASE_API_URL,
    LOGIN_URL
  } from './constants'

const APICall = (url, headers, convertToJson = true) => {
    const client = {
        method: 'GET',
        headers: getHeaders(),
        body: JSON.stringify(body)
    }
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

const LoginAPI = () => {
    const url = BASE_API_URL + LOGIN_URL
    const client = getClient('POST')
    return APICall(url, client)
}
