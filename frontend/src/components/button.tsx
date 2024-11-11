import React from "react";
import { Button } from "antd";

import { injectNavigate } from "../utils";

interface NeumorphismButtonProps {
    icon: React.ReactNode;
    title: string;
    path: string;
}

class NeumorphismButton extends React.Component<
    NeumorphismButtonProps & { navigate: (_: string) => void }
> {
    handleClicked = () => {
        this.props.navigate(this.props.path);
    };

    render() {
        const { icon, title } = this.props;
        return (
            <Button type="primary" className="neumorphism-button" onClick={this.handleClicked}>
                <span className="icon">{icon}</span>
                <span className="text">{title}</span>
            </Button>
        );
    }
}

export default injectNavigate<NeumorphismButtonProps>(NeumorphismButton);
