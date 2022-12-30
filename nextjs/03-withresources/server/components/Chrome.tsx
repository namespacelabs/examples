// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import { ReactNode } from "react";
import Image from "next/image";

export function Chrome(props: { description: ReactNode; children: ReactNode }) {
	return (
		<div>
			<div className="flex flex-col absolute top-0 left-0 right-0 bottom-0 justify-center items-center">
				<div>
					<div className="mb-8 flex">
						<Image src="/logo.svg" width={48} height={48} className="w-12 h-12" alt="logo" />
					</div>

					<div className="mb-24">
						<div className="mx-auto max-w-2xl text-left">
							<h2 className="text-xl font-bold tracking-tight text-white sm:text-4xl">
								<span className="block">Welcome to Namespace.</span>
							</h2>
							<div className="mt-4 text-lg leading-6 text-indigo-200">{props.description}</div>
							<a
								href="https://namespace.so/docs"
								target="_blank"
								rel="noreferrer"
								className="inline-flex items-center justify-center text-base font-medium text-indigo-600">
								Documentation
								<svg
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
									strokeWidth="1.5"
									stroke="currentColor"
									className="w-5 h-5 ml-2">
									<path
										strokeLinecap="round"
										strokeLinejoin="round"
										d="M13.5 6H5.25A2.25 2.25 0 003 8.25v10.5A2.25 2.25 0 005.25 21h10.5A2.25 2.25 0 0018 18.75V10.5m-10.5 6L21 3m0 0h-5.25M21 3v5.25"
									/>
								</svg>
							</a>
						</div>
					</div>
					<div className="mx-auto max-w-2xl text-left rounded-xl border border-indigo-200 overflow-hidden">
						{props.children}
					</div>
				</div>
			</div>
		</div>
	);
}
