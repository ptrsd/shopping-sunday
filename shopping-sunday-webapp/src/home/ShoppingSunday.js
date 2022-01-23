import React, {useEffect, useState} from "react";
import {ShoppingSundayPanel} from "./ShoppingSundayPanel";

export const ShoppingSunday = (props) => {
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

    return (
        <ShoppingSundayPanel details={reasons} isShoppingSunday={isShoppingSunday}/>
    )
}
