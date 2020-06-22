import {
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
    const url = SPONSORS_URL
    const client = getClient('GET')
    return APICall(url, client)
}

const mailingAPI = (url, name, email, body) => {
    let formData = new FormData();
    formData.append('name', name);
    formData.append('email', email);
    formData.append('body', body);

    const options = {
        method: 'POST',
        body: formData
    }

    return APICall(url, options, false)
}

const fetchSocials = () => {
    const url = SOCIAL_URL
    const options = {
        method: 'GET',
    }
    return APICall(url, options)
}

const fetchFaqs = () => {
    const url = FAQ_URL
    const options = {
        method: 'GET',
    }
    return APICall(url, options)
}

const APIClient = {
    sponsorsAPI,
    mailingAPI,
    fetchSocials,
    fetchFaqs
}
export default APIClient