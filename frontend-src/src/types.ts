enum SecretType {
    UploadFileType,
    UploadStringType,
    DownloadFileType,
    DownloadStringType,
}

enum DownloadStatus {
    Available,
	Download,
	Decryption,
    Done,
}

enum UploadStatus {
    Encryption,
    Upload,
    Done,
}

interface UploadFile {
    type: UploadFileType;
    id: string;
    name: string;
    date: Date;
    speed: number;
    progress: number;
    status: UploadStatus;
}

interface UploadString {
    type: UploadStringType;
    id: string;
    content: string;
    date: Date;
    status: UploadStatus,
}

interface DownloadFile {
    type: DownloadFileType,
    id: string;
    name: string;
    date: Date;
    speed: number;
    progress: number;
    status: DownloadStatus;
}

interface DownloadString {
    type: DownloadStringType,
    id: string;
    content: string;
    date: Date;
    sattus: DownloadStatus;
}

export {
    DownloadStatus,
    UploadStatus,
    DownloadFile,
    DownloadString,
    UploadFile,
    UploadString,
}
