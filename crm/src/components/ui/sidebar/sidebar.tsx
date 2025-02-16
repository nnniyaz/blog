import classes from "./sidebar.module.scss";

export const Sidebar = () => {

    return (
        <div className={classes.sidebar}>
            <div className={classes.sidebar__header}>
                <div className={classes.sidebar__header__title}>
                    Blog
                </div>
            </div>
            <div className={classes.sidebar__body}>
                <div className={classes.sidebar__body__list}>
                    <div className={classes.sidebar__body__list__item}>
                        Bio
                    </div>
                    <div className={classes.sidebar__body__list__item}>
                        Articles
                    </div>
                    <div className={classes.sidebar__body__list__item}>
                        Projects
                    </div>
                    <div className={classes.sidebar__body__list__item}>
                        Latest Changes
                    </div>
                </div>
            </div>
        </div>
    )
}
