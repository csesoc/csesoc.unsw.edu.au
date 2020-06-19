export const SPONSORS_URL = '/api/v1/sponsors'

export const GENERAL_FORM_URL = '/api/v1/mailing/general'
export const SPONSORSHIP_FORM_URL = '/api/v1/mailing/sponsorship'
export const FEEDBACK_FORM_URL = '/api/v1/mailing/feedback'

export const MAILING_URL = {
   "general": GENERAL_FORM_URL,
   "sponsorship": SPONSORSHIP_FORM_URL,
   "feedback": FEEDBACK_FORM_URL
}
Object.freeze(MAILING_URL)
