import React from "react";

import { UploadOutlined } from "@ant-design/icons";
import { Button, message, Upload } from "antd";
import SparkMD5 from "spark-md5";

interface UploadBlockState {
    uploading: boolean;
    progess: number; // [0, 1]
}

async function hashFile(file: File): Promise<string> {
    const chunkSize = 2 * 1024 * 1024; // 2MB chunks
    const totalChunks = Math.ceil(file.size / chunkSize);
    const spark = new SparkMD5.ArrayBuffer();
    let res = "";

    let currentChunk = 0;

    const fileReader = new FileReader();

    fileReader.onload = (event) => {
        if (event.target?.result instanceof ArrayBuffer) {
            spark.append(event.target.result);
        }

        currentChunk++;
        if (currentChunk < totalChunks) {
            loadNextChunk();
        } else {
            res = spark.end();
        }
    };

    fileReader.onerror = () => {
        throw new Error("Failed to read file");
    };

    function loadNextChunk() {
        const start = currentChunk * chunkSize;
        const end = Math.min(start + chunkSize, file.size);
        fileReader.readAsArrayBuffer(file.slice(start, end));
    }

    loadNextChunk();

    return res;
}

class UploadBlock extends React.Component<{}, UploadBlockState> {
    constructor(props: {}) {
        super(props);

        this.state = {
            uploading: false,
            progess: 0,
        };
    }

    handleUpload = async (options: any) => {
        // The slice operation is not supported by UploadFile, use the raw file instead
        const file = options.file;
        if (!file) {
            message.error("Unable to get raw file");
            return;
        }
        const key: string = await hashFile(file);

        const chunkSize = 10 * 1024 * 1024; // 10MB
        const totalChunks = Math.ceil((file.size ?? 0) / chunkSize);

        this.setState({
            uploading: true,
        });

        const promises: Promise<void>[] = [];

        for (let i = 0; i < totalChunks; i++) {
            const start = i * chunkSize;
            const end = Math.min(start + chunkSize, file.size ?? 0);
            const chunk = file.slice(start, end);

            // Construct a multipart form data
            const formData = new FormData();

            formData.append("chunk", chunk, file.name);
            formData.append("index", i.toString());
            formData.append("total", totalChunks.toString());
            formData.append("key", key);

            promises.push(
                (async () => {
                    await fetch("/api/upload", {
                        method: "POST",
                        body: formData,
                    });
                    // Set progress
                    this.setState({
                        progess: this.state.progess + 1 / totalChunks,
                    });
                })()
            );
        }

        Promise.all(promises)
            .then((responses) => {
                // Merge chunks
            })

            .catch((error) => {
                // Cleanup
            });
    };

    render() {
        // Render current state
        const { uploading } = this.state;
        return (
            <>
                <Upload
                    name="file"
                    multiple={false}
                    customRequest={this.handleUpload}
                >
                    <Button icon={<UploadOutlined />}>Select File</Button>
                    {uploading && <p>Uploading...</p>}
                </Upload>
            </>
        );
    }
}

export default UploadBlock;
