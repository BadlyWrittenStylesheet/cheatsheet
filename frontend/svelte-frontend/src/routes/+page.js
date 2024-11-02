export const load = async ({ fetch }) => {
    const res = await fetch("http://localhost:55003/cheatsheets");
    const cheatsheets = await res.json();
    return { cheatsheets };
}
// export const prerender = true;

