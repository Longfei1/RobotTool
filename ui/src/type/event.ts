export interface EvOpenJsonEidtor {
    data: object;
    onClose?: (data: object) => void;
    onSave?: (data: object) => void;
}