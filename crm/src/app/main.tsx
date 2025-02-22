import { AppRoutes } from '@routes/index'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from '@app/store'

import './main.scss'
import { BrowserRouter } from 'react-router-dom'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <BrowserRouter>
                <AppRoutes />
            </BrowserRouter>
        </Provider>
    </StrictMode>,
)
