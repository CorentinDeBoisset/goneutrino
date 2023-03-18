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
  id: string;
  name: string;
  date: Date;
  size: number;
  speed: number;
  progress: number;
  status: UploadStatus;
}

interface UploadString {
  id: string;
  content: string;
  date: Date;
  status: UploadStatus;
}

interface DownloadFile {
  id: string;
  name: string;
  date: Date;
  size: number;
  speed: number;
  progress: number;
  status: DownloadStatus;
}

interface DownloadString {
  id: string;
  content: string;
  date: Date;
  status: DownloadStatus;
}

interface SecretItem {
  id: string;
  downloadString?: DownloadString;
  downloadFile?: DownloadFile;
  uploadString?: UploadString;
  uploadFile?: UploadFile;
}

export {
  DownloadStatus,
  UploadStatus,
  DownloadFile,
  DownloadString,
  UploadFile,
  UploadString,
  SecretItem,
};
