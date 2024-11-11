import React from "react";
import { useNavigate } from "react-router-dom";

// P preserves the original props definition
export function injectNavigate<P>(
    Component: React.ComponentType<P & { navigate: (_: string) => void }>
) {
    return function (props: P) {
        const navigate = useNavigate();
        return <Component {...props} navigate={navigate} />;
    };
}
