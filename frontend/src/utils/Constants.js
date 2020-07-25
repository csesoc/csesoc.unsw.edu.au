export const SPONSORS_URL = '/api/v1/sponsors'
export const SOCIAL_URL = '/api/v1/social'
export const FAQ_URL = '/api/v1/faq'
export const EVENTS_URL = '/api/v1/events'
export const RESOURCES_URL = '/api/v1/resources'

export const GENERAL_FORM_URL = '/api/v1/mailing/general'
export const SPONSORSHIP_FORM_URL = '/api/v1/mailing/sponsorship'
export const FEEDBACK_FORM_URL = '/api/v1/mailing/feedback'

export const MAILING_URL = {
   "general": GENERAL_FORM_URL,
   "sponsorship": SPONSORSHIP_FORM_URL,
   "feedback": FEEDBACK_FORM_URL
}
Object.freeze(MAILING_URL)
