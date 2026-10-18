import { z } from "zod";

const passageSchema = z.object({
    id: z.string().optional(),
    header: z.string().trim().optional(),
    verse: z.string().trim().optional(),
    message: z.string().trim().optional(),
    type: z.string().trim().min(1, "Passage type is required."),
    order: z.number().int().min(0, "Passage order must be zero or greater."),
});

const messageFields = {
    title: z.string().trim().min(1, "Title is required."),
    series: z.string().trim().min(1, "Series is required."),
    preacher: z.string().trim().min(1, "Preacher is required."),
    thumbnail: z.string().trim(),
    date: z.string().trim().min(1, "Date is required."),
    passages: z.array(passageSchema),
};

export const createMessageSchema = z.object(messageFields);

export const updateMessageSchema = createMessageSchema.extend({
    id: z.string().trim().min(1, "Message ID is required."),
    published: z.boolean(),
});
