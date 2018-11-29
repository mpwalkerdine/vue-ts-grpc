export interface RootState {
    examples: Example[];
}

export interface Example {
    name: string;
    text: string;
}