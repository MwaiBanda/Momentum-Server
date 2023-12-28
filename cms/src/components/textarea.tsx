import { Label, Textarea } from "flowbite-react"
import { Ref } from "react"

export function TextArea({ reference, value, title, placeholder, id, onChange }: { 
    reference?: Ref<HTMLTextAreaElement>,
    title: string, 
    placeholder: string, 
    id: string,
    value?: string
    onChange?: React.ChangeEventHandler<HTMLTextAreaElement> | undefined
}) {
    return (
        <div>
            <div className="mb-2 block">
                <Label htmlFor={id} value={title} />
            </div>
            <Textarea id={id} ref={reference} placeholder={placeholder} value={value} onChange={onChange} required />
        </div>
    )
}