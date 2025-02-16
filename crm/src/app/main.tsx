import { AppRoutes } from '@routes/index'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from '@app/store'

import './main.scss'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Provider store={store}>
            <AppRoutes />
        </Provider>
    </StrictMode>,
)
