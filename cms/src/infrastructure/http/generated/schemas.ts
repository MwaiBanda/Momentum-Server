import { z } from "zod";

const event_Attributes = z
  .object({
    all_day_event: z.boolean(),
    ends_at: z.string(),
    image_url: z.string(),
    name: z.string(),
    open_graph_description: z.string(),
    recurrence_description: z.string(),
    recurring: z.boolean(),
    starts_at: z.string(),
  })
  .passthrough();
const event_Event = z
  .object({ attributes: event_Attributes, id: z.string() })
  .passthrough();
const event_Response = z.object({ data: z.array(event_Event) }).passthrough();
const http_ErrorResponse = z
  .object({ code: z.number().int(), message: z.string() })
  .passthrough();
const meal_UpdateEntryResponse = z
  .object({ id: z.string(), message: z.string() })
  .passthrough();
const meal_UserResponse = z
  .object({
    createdOn: z.string(),
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    phone: z.string(),
  })
  .passthrough();
const meal_VolunteeredMealResponse = z
  .object({
    created_on: z.string(),
    date: z.string(),
    description: z.string(),
    id: z.string(),
    notes: z.string(),
    updates: z.array(meal_UpdateEntryResponse),
    user: meal_UserResponse,
  })
  .passthrough();
const meal_ParticipantResponse = z
  .object({ id: z.string(), user: meal_UserResponse })
  .passthrough();
const meal_Response = z
  .object({
    allergies: z.string(),
    city: z.string(),
    email: z.string(),
    favourites: z.string(),
    id: z.string(),
    instructions: z.string(),
    least_favourites: z.string(),
    meals: z.array(meal_VolunteeredMealResponse),
    num_of_adults: z.number().int(),
    number_of_kids: z.number().int(),
    participants: z.array(meal_ParticipantResponse),
    phone: z.string(),
    preferred_time: z.string(),
    reason: z.string(),
    recipient: z.string(),
    street: z.string(),
    user: meal_UserResponse,
  })
  .passthrough();
const meal_MealVolunteeredRequest = z
  .object({
    date: z.string(),
    description: z.string().optional(),
    meal_id: z.string().optional(),
    notes: z.string().optional(),
    user_id: z.string().optional(),
  })
  .passthrough();
const meal_Input = z
  .object({
    allergies: z.string(),
    city: z.string(),
    email: z.string(),
    favourites: z.string(),
    id: z.string().optional(),
    instructions: z.string(),
    least_favourites: z.string(),
    meals: z.array(meal_MealVolunteeredRequest),
    num_of_adults: z.number().int(),
    number_of_kids: z.number().int(),
    phone: z.string(),
    preferred_time: z.string(),
    reason: z.string(),
    recipient: z.string(),
    street: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const meal_CreateResponse = z
  .object({
    allergies: z.string(),
    city: z.string(),
    email: z.string(),
    favourites: z.string(),
    id: z.string(),
    instructions: z.string(),
    least_favourites: z.string(),
    meals: z.array(meal_MealVolunteeredRequest),
    num_of_adults: z.number().int(),
    number_of_kids: z.number().int(),
    phone: z.string(),
    preferred_time: z.string(),
    reason: z.string(),
    recipient: z.string(),
    street: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const meal_UpdateResponse = z
  .object({
    allergies: z.string(),
    city: z.string(),
    email: z.string(),
    favourites: z.string(),
    id: z.string(),
    instructions: z.string(),
    least_favourites: z.string(),
    meals: z.array(meal_MealVolunteeredRequest),
    num_of_adults: z.number().int(),
    number_of_kids: z.number().int(),
    phone: z.string(),
    preferred_time: z.string(),
    reason: z.string(),
    recipient: z.string(),
    street: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const meal_UserInput = z.object({ id: z.string() }).passthrough();
const meal_VolunteeredMealInput = z
  .object({
    created_on: z.string().optional(),
    date: z.string().optional(),
    description: z.string().optional(),
    id: z.string(),
    notes: z.string().optional(),
    updates: z.array(meal_UpdateEntryResponse).optional(),
    user: meal_UserInput,
  })
  .passthrough();
const meal_VolunteerAssignmentInput = z
  .object({ meal_id: z.string(), volunteered_meal: meal_VolunteeredMealInput })
  .passthrough();
const meal_VolunteerAssignmentResponse = z
  .object({
    meal_id: z.string(),
    volunteered_meal: meal_VolunteeredMealResponse,
  })
  .passthrough();
const meal_UpdateVolunteeredMealInput = z
  .object({ description: z.string(), id: z.string(), notes: z.string() })
  .passthrough();
const meal_UpdateVolunteeredMealResponse = z
  .object({ description: z.string(), id: z.string(), notes: z.string() })
  .passthrough();
const message_CreatePassageInput = z
  .object({
    header: z.string().optional(),
    message: z.string().optional(),
    order: z.number().int(),
    type: z.string(),
    verse: z.string().optional(),
  })
  .passthrough();
const message_CreateInput = z
  .object({
    date: z.string(),
    passages: z.array(message_CreatePassageInput),
    preacher: z.string(),
    series: z.string(),
    thumbnail: z.string(),
    title: z.string(),
  })
  .passthrough();
const message_NoteResponse = z
  .object({ content: z.string(), id: z.string(), userId: z.string() })
  .passthrough();
const message_PassageResponse = z
  .object({
    header: z.string(),
    id: z.string(),
    message: z.string(),
    notes: z.array(message_NoteResponse),
    order: z.number().int(),
    type: z.string(),
    verse: z.string(),
  })
  .passthrough();
const message_MessageResponse = z
  .object({
    createdOn: z.string(),
    date: z.string(),
    hasOrder: z.boolean(),
    id: z.string(),
    passages: z.array(message_PassageResponse),
    preacher: z.string(),
    published: z.boolean(),
    series: z.string(),
    thumbnail: z.string(),
    title: z.string(),
  })
  .passthrough();
const message_UpdatePassageInput = z
  .object({
    header: z.string().optional(),
    id: z.string(),
    message: z.string().optional(),
    verse: z.string().optional(),
  })
  .passthrough();
const message_UpdateInput = z
  .object({
    date: z.string(),
    id: z.string(),
    passages: z.array(message_UpdatePassageInput),
    preacher: z.string(),
    published: z.boolean(),
    series: z.string(),
    thumbnail: z.string(),
    title: z.string(),
  })
  .passthrough();
const message_MessageListResponse = z
  .object({ data: z.array(message_MessageResponse) })
  .passthrough();
const message_CreateNoteInput = z
  .object({ content: z.string(), passageId: z.string(), userId: z.string() })
  .passthrough();
const message_CreateNoteResponse = z
  .object({ content: z.string(), id: z.string(), userId: z.string() })
  .passthrough();
const message_UpdateNoteInput = z
  .object({ content: z.string(), id: z.string() })
  .passthrough();
const notification_Input = z
  .object({ body: z.string(), title: z.string(), topic: z.string() })
  .passthrough();
const notification_Response = z
  .object({ body: z.string(), title: z.string(), topic: z.string() })
  .passthrough();
const payment_Input = z
  .object({
    amount: z.number().int(),
    email: z.string(),
    fullname: z.string(),
    phone: z.string(),
  })
  .passthrough();
const payment_Response = z
  .object({
    customer: z.string(),
    ephemeralKey: z.string(),
    paymentIntent: z.string(),
    publishableKey: z.string(),
  })
  .passthrough();
const sermon_Sermon = z
  .object({
    categories: z.unknown(),
    date: z.object({ carbon: z.string(), date: z.string() }).passthrough(),
    detail_url: z.string(),
    id: z.string(),
    interactions: z
      .object({ likes: z.number().int(), shares: z.number().int() })
      .passthrough(),
    interactive_note: z.unknown(),
    keywords: z.unknown(),
    media: z
      .object({
        audio: z.unknown(),
        embed: z.unknown(),
        image: z.unknown(),
        notes: z.unknown(),
        video: z.string(),
        video_thumbnail: z.string(),
      })
      .passthrough(),
    passages: z.array(z.unknown()),
    preacher: z.string(),
    series: z
      .object({
        description: z.unknown(),
        image: z.unknown(),
        slug: z.string(),
        title: z.string(),
      })
      .passthrough(),
    slug: z.string(),
    summary: z.unknown(),
    text: z.unknown(),
    title: z.string(),
  })
  .passthrough();
const sermon_Response = z
  .object({
    current_page: z.number().int(),
    data: z.array(sermon_Sermon),
    first_page_url: z.string(),
    from: z.number().int(),
    last_page: z.number().int(),
    last_page_url: z.string(),
    links: z.array(
      z
        .object({ active: z.boolean(), label: z.string(), url: z.unknown() })
        .passthrough()
    ),
    next_page_url: z.string(),
    path: z.string(),
    per_page: z.number().int(),
    prev_page_url: z.unknown(),
    to: z.number().int(),
    total: z.string(),
  })
  .passthrough();
const volunteer_service_ServiceResponse = z
  .object({
    id: z.string(),
    image: z.string(),
    name: z.string(),
    type: z.string(),
  })
  .passthrough();
const volunteer_service_CreateDayInput = z
  .object({ date: z.string() })
  .passthrough();
const volunteer_service_CreateInput = z
  .object({
    days: z.array(volunteer_service_CreateDayInput),
    description: z.string(),
    email: z.string(),
    id: z.string().optional(),
    organizer: z.string(),
    phone: z.string(),
    service_id: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const volunteer_service_CreateResponse = z
  .object({
    days: z.array(volunteer_service_CreateDayInput),
    description: z.string(),
    email: z.string(),
    id: z.string(),
    organizer: z.string(),
    phone: z.string(),
    service_id: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const volunteer_service_UserInput = z.object({ id: z.string() }).passthrough();
const volunteer_service_UpdateDayInput = z
  .object({
    date: z.string(),
    id: z.string(),
    notes: z.string(),
    user: volunteer_service_UserInput,
  })
  .passthrough();
const volunteer_service_UserResponse = z
  .object({
    createdOn: z.string(),
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    meals: z.array(z.unknown()),
    phone: z.string(),
  })
  .passthrough();
const volunteer_service_UpdateDayResponse = z
  .object({
    date: z.string(),
    id: z.string(),
    notes: z.string(),
    user: volunteer_service_UserResponse,
  })
  .passthrough();
const volunteer_service_VolunteerServiceDayResponse = z
  .object({
    date: z.string(),
    id: z.string(),
    notes: z.string(),
    user: volunteer_service_UserResponse,
  })
  .passthrough();
const volunteer_service_VolunteerServiceResponse = z
  .object({
    days: z.array(volunteer_service_VolunteerServiceDayResponse),
    description: z.string(),
    email: z.string(),
    id: z.string(),
    organizer: z.string(),
    phone: z.string(),
    service_id: z.string(),
    user_id: z.string(),
  })
  .passthrough();
const volunteer_service_ServiceDetailResponse = z
  .object({
    id: z.string(),
    image: z.string(),
    name: z.string(),
    services: z.array(volunteer_service_VolunteerServiceResponse),
    type: z.string(),
  })
  .passthrough();
const transaction_UserResponse = z
  .object({
    createdOn: z.string(),
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    phone: z.string(),
  })
  .passthrough();
const transaction_Response = z
  .object({
    amount: z.number().int(),
    createdOn: z.string(),
    date: z.string(),
    description: z.string(),
    id: z.string(),
    user: transaction_UserResponse,
  })
  .passthrough();
const transaction_Input = z
  .object({
    amount: z.number().int(),
    date: z.string(),
    description: z.string(),
    email: z.string().optional(),
    fullname: z.string().optional(),
    id: z.string().optional(),
    phone: z.string().optional(),
    userId: z.string().optional(),
  })
  .passthrough();
const user_CreateInput = z
  .object({
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    phone: z.string(),
  })
  .passthrough();
const user_Response = z
  .object({
    createdOn: z.string(),
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    phone: z.string(),
  })
  .passthrough();
const user_UpdateInput = z
  .object({
    email: z.string(),
    fullname: z.string(),
    id: z.string(),
    phone: z.string(),
  })
  .passthrough();

export const schemas = {
  event_Attributes,
  event_Event,
  event_Response,
  http_ErrorResponse,
  meal_UpdateEntryResponse,
  meal_UserResponse,
  meal_VolunteeredMealResponse,
  meal_ParticipantResponse,
  meal_Response,
  meal_MealVolunteeredRequest,
  meal_Input,
  meal_CreateResponse,
  meal_UpdateResponse,
  meal_UserInput,
  meal_VolunteeredMealInput,
  meal_VolunteerAssignmentInput,
  meal_VolunteerAssignmentResponse,
  meal_UpdateVolunteeredMealInput,
  meal_UpdateVolunteeredMealResponse,
  message_CreatePassageInput,
  message_CreateInput,
  message_NoteResponse,
  message_PassageResponse,
  message_MessageResponse,
  message_UpdatePassageInput,
  message_UpdateInput,
  message_MessageListResponse,
  message_CreateNoteInput,
  message_CreateNoteResponse,
  message_UpdateNoteInput,
  notification_Input,
  notification_Response,
  payment_Input,
  payment_Response,
  sermon_Sermon,
  sermon_Response,
  volunteer_service_ServiceResponse,
  volunteer_service_CreateDayInput,
  volunteer_service_CreateInput,
  volunteer_service_CreateResponse,
  volunteer_service_UserInput,
  volunteer_service_UpdateDayInput,
  volunteer_service_UserResponse,
  volunteer_service_UpdateDayResponse,
  volunteer_service_VolunteerServiceDayResponse,
  volunteer_service_VolunteerServiceResponse,
  volunteer_service_ServiceDetailResponse,
  transaction_UserResponse,
  transaction_Response,
  transaction_Input,
  user_CreateInput,
  user_Response,
  user_UpdateInput,
};
