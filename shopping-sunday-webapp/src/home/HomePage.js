import React from "react";
import {useParams} from "react-router-dom";
import {ShoppingSunday} from "./ShoppingSunday";
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
                        <ShoppingSunday date={date ? date : (new Date()).toISOString().split("T")[0]}/>
                    </div>
                </div>
            </div>
        </>
    );
};
