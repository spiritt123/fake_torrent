import React, { useState, useCallback } from 'react';
import { uploadFile } from '../api/Api';

const FileUpload = ({ onSuccess, onCancel }) => {
    const [file, setFile] = useState(null);
    const [isDragging, setIsDragging] = useState(false);
    const [isUploading, setIsUploading] = useState(false);

    const handleDragEnter = useCallback((e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(true);
    }, []);

    const handleDragLeave = useCallback((e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);
    }, []);

    const handleDragOver = useCallback((e) => {
        e.preventDefault();
        e.stopPropagation();
    }, []);

    const handleDrop = useCallback((e) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);

        if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
            setFile(e.dataTransfer.files[0]);
            e.dataTransfer.clearData();
        }
    }, []);

    const handleFileChange = (e) => {
        if (e.target.files && e.target.files.length > 0) {
            setFile(e.target.files[0]);
        }
    };

    const handleSubmit = async () => {
        if (!file) return;

        setIsUploading(true);
        try {
            await uploadFile(file);
            onSuccess();
        } catch (error) {
            console.error('Upload failed:', error);
        } finally {
            setIsUploading(false);
        }
    };

    return (
        <div className="upload-container m-win">
            <h2>Загрузка файла</h2>

            <div
                className={`drop-area ${isDragging ? 'dragging' : ''}`}
                onDragEnter={handleDragEnter}
                onDragLeave={handleDragLeave}
                onDragOver={handleDragOver}
                onDrop={handleDrop}
            >
                {file ? (
                    <div className="file-selected">
                        <p>Выбранный файл: <strong>{file.name}</strong></p>
                        <p>Размер: {file.size}</p>
                    </div>
                ) : (
                    <>
                        <p>Перетащите сюда файл</p>
                        <p>или</p>
                        <label className="file-input-label">
                            Выберите файл
                            <input
                                type="file"
                                onChange={handleFileChange}
                                className="file-input"
                            />
                        </label>
                    </>
                )}
            </div>

            <div className="upload-actions">
                <button
                    onClick={onCancel}
                    className="cancel-button btn"
                >
                    🔙 Вернуться
                </button>
                <button
                    onClick={handleSubmit}
                    disabled={!file || isUploading}
                    className="upload-button btn"
                >
                    {isUploading ? '🕐 Загрузка...' : '💾 Загрузить'}
                </button>
            </div>
        </div>
    );
};

export default FileUpload;