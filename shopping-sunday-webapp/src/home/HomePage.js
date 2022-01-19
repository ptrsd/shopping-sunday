import React from "react";
import {useParams} from "react-router-dom";
import {ShoppingSundayBlock} from "./ShoppingSundayBlock";
import {useTranslation} from "react-i18next";
import "../translations/i18n";

export const HomePage = () => {
    const {t} = useTranslation();
    let {date} = useParams();

    return (
        <>
            <div className="container">
                <div className="row">
                    <h2 className="text-center">{t("title")}</h2>
                    <div className="row">
                        <ShoppingSundayBlock date={date}/>
                    </div>
                </div>
            </div>
        </>
    );
};
