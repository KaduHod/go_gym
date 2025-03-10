## **Prompt: Generate Joint Movements JSON for an Exercise**

You are an expert in biomechanics and kinesiology. Your task is to analyze the given exercise and return a structured JSON object listing the associated joint movements based on the provided exercise name and description.

### **Input Data:**
- **Exercise Name:** `{exercise_name}`
- **Exercise Description:** `{exercise_description}`

### **Additional Information:**
You have access to the following reference data:
- **List of Joints:** `{joints_list}`
- **List of Movements:** `{movements_list}`
- **List of Muscle Portions:** `{muscles_list}`

### **Output Format:**
Return a JSON array where each entry represents a joint movement involved in the exercise, formatted as follows:

```json
[
  {
    "moviment_id": <integer>,
    "moviment_name": "<string>",
    "joint_id": <integer>,
    "joint_name": "<string>",
    "muscle_id": <integer>,
    "muscle_name": "<string>"
  }
]
### Example Output:
For a Pull-Up exercise, the expected output format would be similar to:

```json
[
  {
    "moviment_id": 1,
    "moviment_name": "Flexion",
    "joint_id": 1,
    "joint_name": "Elbow",
    "muscle_id": 1,
    "muscle_name": "Upper major"
  },
  ...
]```
### Important Guidelines:
- Identify all major joint movements involved in the exercise.
- Map each joint movement to its corresponding muscle portion.
- Use the provided reference lists to ensure consistency in naming and IDs.
- Output only valid movements relevant to the exercise mechanics.
- The response should be in pure JSON format, without additional explanations.

