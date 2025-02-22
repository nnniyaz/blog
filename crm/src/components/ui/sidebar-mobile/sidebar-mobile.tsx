import { useState } from 'react'
import { Drawer } from 'antd'

export const SidebarMobile = () => {
    const [collapsed, setCollapsed] = useState(false)

    const toggleCollapsed = () => {
        setCollapsed(!collapsed)
    }

    return <Drawer open={!collapsed} onClose={toggleCollapsed}></Drawer>
}
