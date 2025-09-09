import Link from "next/link";

export default function AdminPage(){
    return (
        <div>
            <h1>Admin Page</h1>
            <div>
                <Link href="/admin/car">Car</Link>
                <Link href="/admin/brand">Brand</Link>
                <Link href="/admin/type">Type</Link>
            </div>
        </div>
    );
}