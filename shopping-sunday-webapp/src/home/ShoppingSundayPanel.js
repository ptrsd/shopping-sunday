import React from "react";
import {useTranslation} from "react-i18next";
import "../translations/i18n";

export const ShoppingSundayPanel = (props) => {
    const {t} = useTranslation();

    const createInfoSection = () => {
        if (props.details.length) {
            return props.details.map((detail, idx) => <p key={idx}>{t(detail.id)}</p>)
        } else {
            return <p key={0}>{t("enjoyShopping")}</p>
        }
    };

    return (
        <>
            <div className="col-md-12 text-center">
                <div className="box">
                    <div className="box-content">
                        <h1 className="tag-title">{props.isShoppingSunday ? t("isShoppingSunday") : t("notShoppingSunday")}</h1>
                        <hr/>
                        {
                            createInfoSection()
                        }
                    </div>
                </div>
            </div>
        </>
    )
}
