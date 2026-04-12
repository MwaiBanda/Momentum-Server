import { Label, TextInput } from "flowbite-react"

export function TextField({ reference, value, title, placeholder, id, onChange }: { 
    reference?: React.RefObject<HTMLInputElement | null>, 
    value?: string,
    title: string, 
    placeholder: string, 
    id: string,
    onChange?: React.ChangeEventHandler<HTMLInputElement> | undefined
 }) {
    return (
        <div>
            <div className="mb-2 block">
                <Label htmlFor={id}>{title}</Label>
            </div>
            <TextInput id={id} ref={reference} placeholder={placeholder} value={value} onChange={onChange} required />
        </div>
    )
}