export const getError = (error) => {
    if (error.isAxiosError && error.response) return error.response.data
    return 'Unexpected error'
}
