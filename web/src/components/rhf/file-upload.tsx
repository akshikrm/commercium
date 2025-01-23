import { Stack, TextField } from "@mui/material"
import { FunctionComponent } from "react"
import { Controller, useFormContext } from "react-hook-form"
import RenderList from "@components/render-list"
import useRHFImageUpload from "./hooks/use-rhf-image-upload"
import Image from "@components/image"

type Props = {
	label: string
	name: string
}

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
	const { control } = useFormContext()
	const { previews, mutation } = useRHFImageUpload(name)
	return (
		<>
			<Stack direction='row'>
				<RenderList
					list={previews}
					render={preview => (
						<Image key={preview.publicID} preview={preview} />
					)}
				/>
			</Stack>
			<Controller
				name={name}
				control={control}
				render={({ fieldState }) => {
					const { error, invalid } = fieldState
					return (
						<TextField
							type='file'
							onChange={async e => {
								const { files } = e.target as HTMLInputElement
								if (files?.length) {
									mutation.mutate(files)
								}
							}}
							label={label}
							error={invalid}
							helperText={error?.message}
						/>
					)
				}}
			/>
		</>
	)
}

export default RHFImageUpload
