const Footer = () => {
    return (
        <>
            <div className="w-[100%] h-auto mt-1 bg-green-600 flex justify-center pt-5 pb-12">
                <div className="grid grid-cols-3 gap-80">
                    <div>
                        <h2 className="text-4xl font-bold text-white">MVTWARE</h2>
                        <ul>
                            <li>
                                JACKET
                            </li>
                            <li>
                                PANTS
                            </li>
                        </ul>
                    </div>
                    <div>
                        <ul>
                            <li>
                                Help
                            </li>
                            <li>
                                Contact Us
                            </li>
                        </ul>
                    </div>
                    <div className="flex justify-center">
                        <ul>
                            Sosmed
                            <li>
                                instagram
                            </li>
                            <li>
                                facebook
                            </li>
                            <li>
                                tiktok
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Footer