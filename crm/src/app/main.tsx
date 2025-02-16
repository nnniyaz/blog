import { AppRoutes } from '@/routes'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from '@app/store'

import './main.css'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <AppRoutes></AppRoutes>
        </Provider>
    </StrictMode>,
)
