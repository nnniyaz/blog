import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import {Provider} from 'react-redux'

import {store} from './store'
import './main.scss'
import {AppRoutes} from '../routes'

createRoot(document.getElementById('root')!).render(
    <Provider store={store}>
        <StrictMode>
            <AppRoutes>

            </AppRoutes>
        </StrictMode>
    </Provider>,
)
