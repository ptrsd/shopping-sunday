import React from "react";
import {useParams} from "react-router-dom";
import {ShoppingSundayBlock} from "./ShoppingSundayBlock";

export const HomePage = () => {
    let {date} = useParams();

    return (
        <>
            <div className="container">
                <div className="row">
                    <h2 className="text-center">Handlowa niedziela?</h2>
                    <div className="row">
                        <ShoppingSundayBlock date={date} />
                    </div>
                </div>
            </div>
        </>
    );
};
