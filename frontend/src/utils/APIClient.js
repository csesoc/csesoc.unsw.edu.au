import {
  SPONSORS_URL, EVENTS_URL, SOCIAL_URL, FAQ_URL, RESOURCES_URL
} from './Constants';

const APICall = (url, headers, convertToJson = true) => fetch(url, headers)
  .then((res) => (convertToJson ? res.json() : res))
  .then((res) => res);

const getHeaders = (token) => {
  const headers = {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  };
  if (token) headers.Authorization = `Bearer ${token}`;
  return headers;
};

const getClient = (method, body) => {
  const client = {
    method,
    headers: getHeaders()
  };
  if (body) client.body = body;
  return client;
};

const fetchSponsors = () => {
  const url = SPONSORS_URL;
  const client = getClient('GET');
  return APICall(url, client);
};

const eventsAPI = () => {
  const url = EVENTS_URL;
  const client = getClient('GET');
  return APICall(url, client);
};

const mailingAPI = (url, name, email, body) => {
  const formData = new FormData();
  formData.append('name', name);
  formData.append('email', email);
  formData.append('body', body);

  const options = {
    method: 'POST',
    body: formData
  };

  return APICall(url, options, false);
};

const socialsAPI = () => {
  const url = SOCIAL_URL;
  const client = getClient('GET');
  return APICall(url, client);
};

const faqsAPI = () => {
  const url = FAQ_URL;
  const client = getClient('GET');
  return APICall(url, client);
};

const resourcesAPI = (param) => {
  const url = RESOURCES_URL + param;
  const client = getClient('GET');
  return APICall(url, client);
};

const APIClient = {
  fetchSponsors,
  mailingAPI,
  eventsAPI,
  socialsAPI,
  faqsAPI,
  resourcesAPI
};
export default APIClient;
