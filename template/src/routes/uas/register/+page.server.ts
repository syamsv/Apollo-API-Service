import type { Actions } from './$types';
import { env } from '$env/dynamic/private'
export const actions = {
    default: async ({ request,fetch }) => {
        const data = await request.formData()
        const email = data.get('email')
        const password = data.get('password')
        const firstname = data.get('firstname')
        const lastname = data.get('lastname')
        
        const response = await fetch(`http://${env.SERVER_HOST}:${env.SERVER_PORT}/api/register`,{
            method:'POST',
            body:JSON.stringify({
                email,
                password,
                firstname,
                lastname
              }),
              headers: {"Content-Type": "application/json"}
        })
        console.log(await response.json())
    }
} satisfies Actions;