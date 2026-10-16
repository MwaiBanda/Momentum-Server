import { Label, Textarea } from "flowbite-react"

export function TextArea({ reference, value, title, placeholder, id, onChange }: { 
    reference?: React.RefObject<HTMLTextAreaElement | null>,
    title: string, 
    placeholder: string, 
    id: string,
    value?: string
    onChange?: React.ChangeEventHandler<HTMLTextAreaElement> | undefined
}) {
    return (
        <div>
            <div className="mb-2 block">
                <Label htmlFor={id}>{title}</Label>
            </div>
            <Textarea id={id} ref={reference} placeholder={placeholder} value={value} onChange={onChange} required rows={6}/>
        </div>
    )
}