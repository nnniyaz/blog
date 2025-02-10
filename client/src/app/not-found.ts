import { headers } from 'next/headers'
import { redirect } from 'next/navigation'

export default async function NotFound() {
    const headersList = await headers()
    redirect(`/${headersList.get('x-current-path')?.split('/')[1] || ''}`)
}
