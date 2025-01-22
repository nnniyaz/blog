import {Outlet} from 'react-router-dom'
import {Routes} from "../../domain/base/routes/routes";
import {useTypedSelector} from "../../lib/hooks/useTypedSelector";
import {useActions} from "../../lib/hooks/useActions";
import {Lang} from "../../app/store/reducers/system/systemSlice";

export const Layout = () => {
    const {lang} = useTypedSelector(state => state.system);
    const {changeLang} = useActions();
    return (
        <div>
            <h1>Layout</h1>
            <h2>Current language: {lang}</h2>
            <select onChange={(event) => changeLang(event.target.value)}>
                <option defaultChecked={true} value={lang}>{lang}</option>
                <option value={Lang.EN}>{Lang.EN}</option>
                <option value={Lang.KZ}>{Lang.KZ}</option>
                <option value={Lang.RU}>{Lang.RU}</option>
            </select>
            <ol>
                <li><a href={Routes.HOME}>HOME</a></li>
                <li><a href={Routes.USER}>USER</a></li>
                <li><a href={Routes.AUTHOR}>AUTHOR</a></li>
                <li><a href={Routes.PROJECT}>PROJECT</a></li>
                <li><a href={Routes.ARTICLE}>ARTICLE</a></li>
                <li><a href={Routes.BOOK}>BOOK</a></li>
                <li><a href={Routes.SETTINGS}>SETTINGS</a></li>
            </ol>
            <Outlet/>
        </div>
    )
}
