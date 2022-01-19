import React, {useEffect, useState} from "react";

export const ShoppingSundayBlock = (props) => {
    let [isShoppingSunday, setIsShoppingSunday] = useState(false);
    let [reasons, setReasons] = useState([]);

    useEffect(() => {
        checkIfShoppingSunday(props.date);
    }, [props.date]);

    const checkIfShoppingSunday = async (d) => {
        await fetch("http://localhost:8080/sunday/" + d)
            .then(res => res.json())
            .then(result => {
                setIsShoppingSunday(result.isShoppingSunday);
                if (Array.isArray(result.reasons) && result.reasons.length) {
                    setReasons(result.reasons)
                }
            });
    };

    const buildReasonSection = () => {
        if (reasons.length) {
            return reasons.map(reason => <p>{reason}</p>)
        } else {
            return <p>Ciesz się zakupami!</p>
        }
    };

    return (
        <>
            <div className="col-md-12 text-center">
                <div className="box">
                    <div className="box-content">
                        <h1 className="tag-title">{isShoppingSunday ? "Tak" : "Nie"}</h1>
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
