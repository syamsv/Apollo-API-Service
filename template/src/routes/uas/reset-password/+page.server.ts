import type { Actions } from './$types';

export const actions = {
    default: async ({request}) => {
        const data = await request.formData()
        const email = await data.get('email')
        console.log(email)
    }
} satisfies Actions;