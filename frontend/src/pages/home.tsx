import React, { ReactNode, CSSProperties } from "react";
import { Flex } from "antd";
import {
    BookOutlined,
    BarsOutlined,
    FileDoneOutlined,
    KeyOutlined,
} from "@ant-design/icons";

import type { FlexProps } from "antd";

import Button from "../components/button";

interface ButtonData {
    icon: ReactNode;
    title: string;
    path: string;
}

const buttons: ButtonData[] = [
    {
        icon: <FileDoneOutlined></FileDoneOutlined>,
        title: "Files",
        path: "/files",
    },
    {
        icon: <KeyOutlined></KeyOutlined>,
        title: "Keys",
        path: "/keys",
    },
    {
        icon: <BookOutlined></BookOutlined>,
        title: "Passwords",
        path: "/passwords",
    },
];

const boxStyle: CSSProperties = {
    width: "100%",
    height: "100%",
};

class HomePage extends React.Component {
    render() {
        return (
            <div>
                <Flex style={boxStyle} justify="center" align="center">
                    {buttons.map(({ icon, title, path }, _) => {
                        return (
                            <Button
                                icon={icon}
                                title={title}
                                path={path}
                            ></Button>
                        );
                    })}
                </Flex>
            </div>
        );
    }
}

export default HomePage;