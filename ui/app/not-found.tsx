import Image from "next/image";
import Link from "next/link";
import { Button } from "../components/ui/button";

export default function NotFound() {
    return (
        <div className="min-h-screen w-full flex flex-col items-center justify-center p-4">
            <div className="max-w-[600px] text-center flex items-center justify-center">
                <Image
                    src="/images/pc_error.png"
                    alt="404 Error Illustration"
                    width={400}
                    height={300}
                    className="w-full h-auto mb-6"
                    priority
                />
                <div>
                    <h1 className="text-4xl font-bold mb-3">Oops!</h1>
                    <p className="text-muted-foreground text-lg mb-6">
                        We couldn&apos;t find the page you were looking for
                    </p>
                    <Button asChild>
                        <Link href="/">Go home</Link>
                    </Button>
                </div>
            </div>
        </div>
    );
}
