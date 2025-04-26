
import React from 'react';

const FileList = ({ files, isLoading, onFileClick }) => {
    if (isLoading) {
        return <div className="loading">Загрузка файлов...</div>;
    }

    if (!files) {
        return <div className="no-files">Не удалось загрузить файлы </div>;
    }

    // if (files.length === 0) {
    //     return <div className="no-files">Нет доступных файлов</div>;

    // }


    return (
        <div className='file-list-container'>
            <div className="files-count">Found {files.length} files</div>
            <div className="file-list">
                {files.length == 0 && (
                    <div className="no-files">📂 Нет доступных файлов</div>
                )}
                {files.map((file) => (
                    <div
                        key={file.id}
                        className="file-item"
                        onClick={() => onFileClick(file.id)}
                    >
                        <span className="file-name">{file.name}</span>
                        <span className="file-size">{formatFileSize(file.size)}</span>
                    </div>
                ))}
            </div>
        </div>
    );
};

const formatFileSize = (bytes) => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
};



export default FileList;