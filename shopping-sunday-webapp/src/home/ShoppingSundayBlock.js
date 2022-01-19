import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import "../translations/i18n";

export const ShoppingSundayBlock = (props) => {
    const {t} = useTranslation();
    let [isShoppingSunday, setIsShoppingSunday] = useState(false);
    let [reasons, setReasons] = useState([]);

    useEffect(() => {
        checkIfShoppingSunday(props.date);
    }, [props.date]);

    const checkIfShoppingSunday = async (d) => {
        await fetch("http://localhost:8080/sunday/" + d)
            .then(response => {
                if (!response.ok) {
                    return Promise.reject(response.status);
                }
                return response.json();
            })
            .then(result => {
                setIsShoppingSunday(result.isShoppingSunday);
                if (Array.isArray(result.reasons) && result.reasons.length) {
                    setReasons(result.reasons)
                }
            })
            .catch(err => {
                setIsShoppingSunday(false);
                if (err === 422) {
                    setReasons([{id: "incorrectFormatError"}])
                } else {
                    setReasons([{id: "unknownError"}])
                }
            })
    };

    const buildReasonSection = () => {
        if (reasons.length) {
            return reasons.map((reason, idx) => <p key={idx}>{t(reason.id)}</p>)
        } else {
            return <p key={0}>{t("enjoyShopping")}</p>
        }
    };

    return (
        <>
            <div className="col-md-12 text-center">
                <div className="box">
                    <div className="box-content">
                        <h1 className="tag-title">{isShoppingSunday ? t("isShoppingSunday") : t("notShoppingSunday")}</h1>
                        <hr/>
                        {
                            buildReasonSection()
                        }
                    </div>
                </div>
            </div>
        </>
    )
}
