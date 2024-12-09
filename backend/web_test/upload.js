document.getElementById('uploadButton').addEventListener('click', async () => {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];

    if (!file) {
        alert('请选择一个文件');
        return;
    }

    const chunkSize = 1024 * 1024; // 每个分片的大小，这里设置为1MB
    const totalChunks = Math.ceil(file.size / chunkSize);
    const fileName = file.name;

    let uploadedChunks = 0;

    for (let i = 0; i < totalChunks; i++) {
        const start = i * chunkSize;
        const end = Math.min(file.size, start + chunkSize);
        const chunk = file.slice(start, end);

        const formData = new FormData();
        formData.append('file', chunk, fileName);
        formData.append('chunkIndex', i);
        formData.append('totalChunks', totalChunks);

        try {
            const response = await fetch('/upload', {
                method: 'POST',
                body: formData
            });

            if (!response.ok) {
                throw new Error('分片上传失败');
            }

            uploadedChunks++;
            document.getElementById('status').innerText = `上传进度: ${uploadedChunks}/${totalChunks}`;
        } catch (error) {
            console.error('分片上传失败:', error);
            alert('分片上传失败，请重试');
            return;
        }
    }

    // 所有分片上传完成后，调用合并接口
    try {
        const mergeResponse = await fetch('/merge', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ fileName })
        });

        if (!mergeResponse.ok) {
            throw new Error('合并文件失败');
        }

        alert('文件上传成功');
        document.getElementById('status').innerText = '文件上传成功';
    } catch (error) {
        console.error('合并文件失败:', error);
        alert('合并文件失败，请重试');
    }
});